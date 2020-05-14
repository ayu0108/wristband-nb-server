function loadReciveData(content) {
    content = content.replace(/[\[\]\{]+/g, "")
    let arr = content.split("} "); //分割資料
    arr[arr.length - 1] = arr[arr.length - 1].replace(/\}+/g, "");

    // 變成物件後載入資料
    items = [];
    for (i = 0; i < arr.length; i++) {
        data = arr[i].toString().split(" ")
        items.push({
            id: data[0],
            name: data[1],
            addr: data[2],
            mac: data[3],
            data: data[4],
            date: data[5],
            conn_status: data[6],
        });
    }

    var $table = $('#table')
    var $remove = $('#remove')
    var selections = []

    function getIdSelections() {
        return $.map($table.bootstrapTable('getSelections'), function (row) {
            return row.id
        })
    }

    function responseHandler(res) {
        $.each(res.rows, function (i, row) {
            row.state = $.inArray(row.id, selections) !== -1
        })
        return res
    }

    function detailFormatter(index, row) {
        var html = []
        $.each(row, function (key, value) {
            html.push('<p><b>' + key + ':</b> ' + value + '</p>')
        })
        return html.join('')
    }

    function operateFormatter(value, row, index) {
        return [
            '<a class="like" href="javascript:void(0)" title="Like">',
            '<i class="fa fa-heart"></i>',
            '</a>  ',
            '<a class="remove" href="javascript:void(0)" title="Remove">',
            '<i class="fa fa-trash"></i>',
            '</a>'
        ].join('')
    }

    window.operateEvents = {
        'click .like': function (e, value, row, index) {
            alert('You click like action, row: ' + JSON.stringify(row))
        },
        'click .remove': function (e, value, row, index) {
            $table.bootstrapTable('remove', {
                field: 'id',
                values: [row.id]
            })
        }
    }

    function initTable() {
        $table.bootstrapTable('destroy').bootstrapTable({
            data: items,
            height: 550,
            locale: $('#locale').val(),
            columns: [
                [{
                    title: 'ID',
                    field: 'id',
                    rowspan: 1,
                    align: 'center',
                    valign: 'middle',
                    sortable: true,
                }, {
                    title: 'name',
                    field: 'name',
                    rowspan: 1,
                    align: 'center',
                    valign: 'middle',
                }, {
                    title: 'addr',
                    field: 'addr',
                    rowspan: 1,
                    align: 'center',
                    valign: 'middle',
                }, {
                    title: 'mac',
                    field: 'mac',
                    rowspan: 1,
                    align: 'center',
                    valign: 'middle',
                }, {
                    title: 'data',
                    field: 'data',
                    rowspan: 1,
                    align: 'center',
                    valign: 'middle',
                }, {
                    title: 'date',
                    field: 'date',
                    rowspan: 1,
                    align: 'center',
                    valign: 'middle',
                    sortable: true,
                }, {
                    title: 'connStatus',
                    field: 'conn_status',
                    rowspan: 1,
                    align: 'center',
                    valign: 'middle',
                }]
            ]
        })
        $table.on('check.bs.table uncheck.bs.table ' +
            'check-all.bs.table uncheck-all.bs.table',
            function () {
                $remove.prop('disabled', !$table.bootstrapTable('getSelections').length)

                // save your data, here just save the current page
                selections = getIdSelections()
                // push or splice the selections if you want to save all data selections
            })
    }

    $(function () {
        initTable()
    })
}