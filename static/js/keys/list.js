layui.use(["element", "table", "form", "okLayer", "okUtils", "okMock","jquery"], function (){
    let table = layui.table;
    let util = layui.util;
    let okLayer = layui.okLayer;
    let okUtils = layui.okUtils;
    $ = layui.jquery;

    util.fixbar({});

    let keyTable = table.render({
        elem: '#tableId',
        url: '/keys/table',
        limit: 20,
        page: true,
        even: true,
        id: 'listReload',
        toolbar: "#toolbarTpl",
        size: "sm",
        cols: [[
            {type: "checkbox", fixed: "left"},
            {field: "ID", title: "ID", width: 0},
            {field: "Pubkey", title: "Public Version", width: 575},
            {field: "Prikey", title: "Private Key", width: 575},
            {field: "State", title: "Status", width: 150},
            {field: "CreateTime", title: "Add Time", width: 150},
            {title: "Operation", width: 100, align: "center", fixed: "right", templet: "#operationTpl"}
        ]],
        done: function (res, curr, count) {
            $(".layui-table-box").find("[data-field='ID']").css("display","none");
        }
    });

    $('#laySearch .layui-btn').on('click', function () {
        table.reload('listReload', {
            where: {
                searchKey: $('#searchKey').val()
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
            case "create":
                create();
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
        }
    });

    function batchDel() {
        okLayer.confirm("Are you sure to bulk delete?", function (index) {
            layer.close(index);
            var idsStr = okUtils.tableBatchCheck(table);
            if (idsStr) {
                okUtils.ajax("/keys/batchDel", "put", {idsStr: idsStr}, true).done(function (response) {
                    console.log(response);
                    okUtils.tableSuccessMsg(response.msg);
                }).fail(function (error) {
                    console.log(error)
                });
            }
        });
    }

    function add() {
        okLayer.open("Add New Key", "add.html", "90%", "90%", null, function () {
            keyTable.reload();
        })
    }

    function create() {
        okLayer.open("Create New Keys", "create.html", "90%", "90%", null, function () {
            keyTable.reload();
        })
    }

    function edit(id) {
        okLayer.open("Edit Device", "edit.html?id=" + id, "90%", "90%", null, function () {
            keyTable.reload();
        })
    }

    function del(id, obj) {
        okLayer.confirm("Are you sure to delete?", function () {
            var jsData = { 'idsStr': id }
            $.post('/keys/ajaxdel', jsData, function (out) {
                if (out.status == 0) {
                    layer.alert('Delete success.', { icon: 1 }, function (index) {
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