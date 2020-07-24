function loadReciveData(content) {
    console.log(content)

    content = content.replace(/map/, '')
    
    content = content.replace(/(\d:{)/g, ',{')
    content = content.replace(/,/, '')
    content = JSON.parse(content)
    console.log(content)

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
            data: content,
            height: 550,
            locale: $('#locale').val(),
            columns: [
                [{
                    title: 'ID',
                    field: 'device_id',
                    rowspan: 1,
                    align: 'center',
                    valign: 'middle',
                    sortable: true,
                }, {
                    title: '接收日期',
                    field: 'CreatedAt',
                    rowspan: 1,
                    align: 'center',
                    valign: 'middle',
                }, {
                    title: '裝置 IP 位址',
                    field: 'ip_address',
                    rowspan: 1,
                    align: 'center',
                    valign: 'middle',
                }, {
                    title: '藍芽名稱',
                    field: 'beacon_name',
                    rowspan: 1,
                    align: 'center',
                    valign: 'middle',
                }, {
                    title: '藍芽資料',
                    field: 'beacon_data',
                    rowspan: 1,
                    align: 'center',
                    valign: 'middle',
                },  {
                    title: '編號',
                    field: 'serial_number',
                    rowspan: 1,
                    align: 'center',
                    valign: 'middle',
                    sortable: true,
                },{
                    title: 'NBIOT 訊號',
                    field: 'singal',
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