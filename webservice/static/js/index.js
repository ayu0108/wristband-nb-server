function loadReciveData(content) {
    content = content.replace(/[\[\]\{]+/g, "")
    let arr = content.split("} "); //分割資料
    arr[arr.length - 1] = arr[arr.length - 1].replace(/\}+/g, "");
    console.log(content)
    console.log(arr)

    data = arr[0].toString().split(" ")
    temp = data[9].split(",")
    console.log(data);
    console.log(temp)

    // 變成物件後載入資料
    items = [];
    for (i = 0; i < arr.length; i++) {
        data = arr[i].toString().split(" ")

        // contentSource: 因為完整資料從伺服器傳過來網頁會跑掉，所以要再抓一次
        contentSource = data[9].split(",") 
        temperature = contentSource[4]
        singal = contentSource[7]
        items.push({
            id: data[0],
            deviceID: data[1],
            date: data[2],
            addr: data[3],
            beaconName: data[4],
            humidity: data[6],
            temperature: temperature,
            serialNumber: data[8],
            deviceStatus: data[7],
            singal: singal,
            content: data[9],
        });
    }

    console.log(items)


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
                    field: 'deviceID',
                    rowspan: 1,
                    align: 'center',
                    valign: 'middle',
                    sortable: true,
                }, {
                    title: '日期',
                    field: 'date',
                    rowspan: 1,
                    align: 'center',
                    valign: 'middle',
                }, {
                    title: 'IP 位址',
                    field: 'addr',
                    rowspan: 1,
                    align: 'center',
                    valign: 'middle',
                }, {
                    title: '藍芽名稱',
                    field: 'beaconName',
                    rowspan: 1,
                    align: 'center',
                    valign: 'middle',
                },{
                    title: '濕度',
                    field: 'humidity',
                    rowspan: 1,
                    align: 'center',
                    valign: 'middle',
                },{
                    title: '溫度',
                    field: 'temperature',
                    rowspan: 1,
                    align: 'center',
                    valign: 'middle',
                },  {
                    title: '編號',
                    field: 'serialNumber',
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
                },{
                    title: '完整資料',
                    field: 'content',
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
}