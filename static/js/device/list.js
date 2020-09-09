layui.use(["element", "table", "form", "laydate", "okLayer", "okUtils", "okMock","jquery"], function (){
    let table = layui.table;
    //let form = layui.form;
    let util = layui.util;
    let laydate = layui.laydate;
    let okLayer = layui.okLayer;
    let okUtils = layui.okUtils;
    //let okMock = layui.okMock;
    $ = layui.jquery;

    util.fixbar({});

    laydate.render({elem: "#startTime", type: "datetime"});
    laydate.render({elem: "#endTime", type: "datetime"});

    //方法级渲染
    let deviceTable = table.render({
        elem: '#tableId',
        url: '/device/table',
        limit: 20,
        page: true,
        even: true,
        id: 'listReload',
        toolbar: "#toolbarTpl",
        size: "sm",
        cols: [[
            {type: "checkbox", fixed: "left"},
            {field: "ID", title: "ID", width: 0},
            {field: "DeviceID", title: "Device ID", width: 200},
            {field: "DeviceName", title: "Device Name", width: 250},
            {field: "DeviceStatus", title: "Status", width: 100},
            {field: "PublicKey", title: "Public Key", width: 500},
            {field: "VersionTitle", title: "Software Version", width: 250},
            {title: "Operation", width: 250, align: "center", fixed: "right", templet: "#operationTpl"}
        ]],
        done: function (res, curr, count) {
            //console.log(res, curr, count)
            $(".layui-table-box").find("[data-field='ID']").css("display","none");
        }
    });

    $('#laySearch .layui-btn').on('click', function () {
        table.reload('listReload', {
            where: {
                startTime: $('#startTime').val(),
                endTime: $('#endTime').val(),
                deviceName: $('#deviceName').val(),
            }
        });
    });

    table.on("toolbar(tableFilter)", function (obj) {
        switch (obj.event) {
            case "batchDel":
                batchDel();
                break;
            case "add":
                add();
                break;
        }
    });

    table.on("tool(tableFilter)", function (obj) {
        let data = obj.data;
        switch (obj.event) {
            case "edit":
                edit(data.ID);
                break;
            case "del":
                del(data.ID, obj);
                break;
            case "chooseKey":
                chooseKey(data.ID);
                break;
            case "chooseSoftware":
                chooseSoftware(data.ID);
                break;
        }
    });

    function chooseKey(id){
        okLayer.open("Choose Key", "choosekey.html?id=" + id, "90%", "90%", null, function () {

            deviceTable.reload();
        });
    }

    function chooseSoftware(id){
        okLayer.open("Choose Software", "choosesoftware.html?id=" + id, "90%", "90%", null, function () {
            deviceTable.reload();
        });
    }

    function batchDel() {
        okLayer.confirm("Are you sure to bulk delete?", function (index) {
            layer.close(index);
            var idsStr = okUtils.tableBatchCheck(table);
            if (idsStr) {
                okUtils.ajax("/device/batchDel", "put", {idsStr: idsStr}, true).done(function (response) {
                    console.log(response);
                    okUtils.tableSuccessMsg(response.msg);
                }).fail(function (error) {
                    console.log(error)
                });
            }
        });
    }

    function add() {
        okLayer.open("Add Device", "add.html", "90%", "90%", null, function () {
            deviceTable.reload();
        })
    }

    function edit(id) {
        okLayer.open("Edit Device", "edit.html?id=" + id, "90%", "90%", null, function () {
            deviceTable.reload();
        })
    }

    function del(id, obj) {
        okLayer.confirm("Are you sure to delete?", function () {
            var jsData = { 'idsStr': id }
            $.post('/device/ajaxdel', jsData, function (out) {
                if (out.status == 0) {
                    layer.alert('Delete Success.', { icon: 1 }, function (index) {
                        layer.close(index);
                        window.location.reload();
                    });
                } else {
                    layer.msg(out.message)
                }
            }, "json");
            obj.del();
        });
    }

});