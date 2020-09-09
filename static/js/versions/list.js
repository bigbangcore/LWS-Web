layui.use(["element", "table", "form", "laydate", "okLayer", "okUtils", "jquery"], function (){
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
    let verTable = table.render({
        elem: '#tableId',
        url: '/versions/table',
        limit: 20,
        page: true,
        even: true,
        id: 'listReload',
        toolbar: "#toolbarTpl",
        size: "sm",
        cols: [[
            {type: "checkbox", fixed: "left"},
            {field: "ID", title: "ID", width: 0},
            {field: "VersionTitle", title: "Title", width: 300},
            {field: "VersionURL", title: "URL", width: 550},
            {field: "VersionSize", title: "Size", width: 250},
            {field: "VersionUpdateTime", title: "Update Time", width: 250},
            {title: "Operation", width: 200, align: "center", fixed: "right", templet: "#operationTpl"}
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
                endTime: $('#endTime').val()
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
            case "push":
                push(data.ID);
                break;
        }
    });

    function batchDel() {
        okLayer.confirm("Are you sure to bulk delete?", function (index) {
            layer.close(index);
            var idsStr = okUtils.tableBatchCheck(table);
            if (idsStr) {
                okUtils.ajax("/versions/batchDel", "put", {idsStr: idsStr}, true).done(function (response) {
                    console.log(response);
                    okUtils.tableSuccessMsg(response.msg);
                }).fail(function (error) {
                    console.log(error)
                });
            }
        });
    }

    function add() {
        okLayer.open("Add Device Version", "add.html", "90%", "90%", null, function () {
            verTable.reload();
        })
    }

    function edit(id) {
        okLayer.open("Edit Device", "edit.html?id=" + id, "90%", "90%", null, function () {
            verTable.reload();
        })
    }

    function push(id){
        okLayer.open("Push Software to Devices", "push.html?id=" + id, "90%", "90%", null, function () {
            verTable.reload();
        })
    }

    function del(id, obj) {
        okLayer.confirm("Are you sure to delete?", function () {
            var jsData = { 'idsStr': id }
            $.post('/versions/ajaxdel', jsData, function (out) {
                if (out.status == 0) {
                    layer.alert('删除成功了', { icon: 1 }, function (index) {
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