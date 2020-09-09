"use strict";
layui.use(["okUtils", "table"],function(){
    var table = layui.table;
    var okUtils = layui.okUtils;
    var $ = layui.jquery;


var option = {
    xAxis: {
        type: 'category',
        data: ['2020-06-29','2020-06-30','2020-07-01','2020-07-02','2020-07-03','2020-07-04','2020-07-05']
    },
    yAxis: {
        type: 'value'
    },
    series: [{
        data: [32, 270, 111, 16, 70, 30, 130],
        type: 'bar'
    }]
};


var devicesSubSourceOption = {
    tooltip: {
        trigger: 'axis'
    },
    toolbox: {
        show: true,
        feature: {
            dataZoom: {
                yAxisIndex: 'none'
            },
            dataView: {readOnly: false},
            magicType: {type: ['line', 'bar']},
            restore: {},
            saveAsImage: {}
        }
    },
    xAxis: {
        type: 'category',
        boundaryGap: false,
        data: ['2020-06-29','2020-06-30','2020-07-01','2020-07-02','2020-07-03']
    },
    yAxis: {
        type: 'value',
        axisLabel: {
            formatter: '{value}'
        }
    },
    series: [
        {
            name: 'dev001',
            type: 'line',
            smooth:true,
            data: [41, 3, 15, 73, 12],
            markPoint: {
                data: [
                    {type: 'max', name: '最大值'},
                    {type: 'min', name: '最小值'}
                ]
            },
            markLine: {
                data: [
                    {type: 'average', name: '平均值'}
                ]
            }
        },
        {
            name: 'dev002',
            type: 'line',
            smooth:true,
            data: [11, 12, 52, 45, 43],
            markPoint: {
                data: [
                    {type: 'max', name: '最大值'},
                    {type: 'min', name: '最小值'}
                ]
            },
            markLine: {
                data: [
                    {type: 'average', name: '平均值'}
                ]
            }
        },
        {
            name: 'dev003',
            type: 'line',
            smooth:true,
            data: [19, 62, 72, 25, 53],
            markPoint: {
                data: [
                    {type: 'max', name: '最大值'},
                    {type: 'min', name: '最小值'}
                ]
            },
            markLine: {
                data: [
                    {type: 'average', name: '平均值'}
                ]
            }
        },
        {
            name: 'dev004',
            type: 'line',
            smooth:true,
            data: [132, 102, 12, 35, 23],
            markPoint: {
                data: [
                    {type: 'max', name: '最大值'},
                    {type: 'min', name: '最小值'}
                ]
            },
            markLine: {
                data: [
                    {type: 'average', name: '平均值'}
                ]
            }
        },
        {
            name: 'dev005',
            type: 'line',
            smooth:true,
            data: [111, 12, 32, 15, 33],
            markPoint: {
                data: [
                    {type: 'max', name: '最大值'},
                    {type: 'min', name: '最小值'}
                ]
            },
            markLine: {
                data: [
                    {type: 'average', name: '平均值'}
                ]
            }
        }
    ]
};

/**
 * 用户访问
 */
function deviceSource() {
    var devicesSub = echarts.init($("#devicesSub")[0], "theme");
    devicesSub.setOption(option);
    okUtils.echartsResize([devicesSub]);
}

deviceSource();

});
