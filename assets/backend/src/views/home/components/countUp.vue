<template>
    <div>
        <p :class="className" :style="{textAlign: 'center', color: color, fontSize: countSize, fontWeight: countWeight}"><span v-cloak :id="idName">{{ startVal }}</span><span>{{ unit }}</span></p>
        <slot name="intro"></slot>
    </div>
</template>

<script>
    import CountUp from 'countup';

    function ForDight(Dight, How) {
        Dight = Math.round(Dight * Math.pow(10, How)) / Math.pow(10, How);
        return Dight;
    }

    function transformValue(val) {
        let endVal = 0;
        let unit = '';
        if (val < 1000) {
            endVal = val;
        } else if (val >= 1000 && val < 1000000) {
            endVal = ForDight(val / 1000, 2);
            unit = 'K+';
        } else if (val >= 1000000 && val < 10000000000) {
            endVal = ForDight(val / 1000000, 2);
            unit = 'M+';
        } else {
            endVal = ForDight(val / 1000000000, 2);
            unit = 'B+';
        }
        return {
            val: endVal,
            unit: unit
        };
    }

    export default {
        data() {
            return {
                unit: '',
                data: {}
            };
        },
        name: 'countUp',
        props: {
            idName: String,
            className: String,
            startVal: {
                type: Number,
                default: 0
            },
            endVal: {
                type: Number,
                required: true
            },
            decimals: {
                type: Number,
                default: 0
            },
            duration: {
                type: Number,
                default: 2
            },
            delay: {
                type: Number,
                default: 500
            },
            options: {
                type: Object,
                default: () => {
                    return {
                        useEasing: true,
                        useGrouping: true,
                        separator: ',',
                        decimal: '.'
                    };
                }
            },
            color: String,
            countSize: {
                type: String,
                default: '30px'
            },
            countWeight: {
                type: Number,
                default: 700
            },
            introText: [String, Number]
        },
        mounted() {
            let that = this;
            that.$nextTick(() => {
                setTimeout(() => {
                    let res = transformValue(that.endVal);
                    let endVal = res.val;
                    that.unit = res.unit;
                    let data = {};
                    that.data = data = new CountUp(that.idName, that.startVal, endVal, that.decimals, that.duration, that.options);
                    if (!data.error) {
                        data.start();
                    }
                }, that.delay);
            });
        },
        watch: {}
    };
</script>
