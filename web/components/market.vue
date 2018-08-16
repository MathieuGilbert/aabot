<template>
    <div class="columns">
        <div class="column is-narrow">
            <div class="cell cell-header">
                &nbsp;
            </div>
            <div class="cell">ETH</div>
            <div class="cell" v-for="token in tokens" :key="token.id">
                {{token.name}}
            </div>
        </div>
        <div class="column is-narrow">
            <div class="cell cell-header">
                Wallet
            </div>
            <div class="cell">{{getBalance("ETH")}}</div>
            <div class="cell" v-for="token in tokens" :key="token.id">
                {{getBalance(token)}}
            </div>
        </div>
        <div class="column is-narrow" v-for="exchange in exchanges" :key="exchange.id">
            <div class="cell cell-header">
                {{exchange.name}}
            </div>
            <div class="cell">{{getMarketBalance(exchange.id, "ETH")}}</div>
            <div class="cell" v-for="token in tokens" :key="token.id">
                {{getMarketBalance(exchange.id, token.id)}}
            </div>
        </div>
    </div>

</template>

<script>
    export default {
        name: "market",
        data() {
            return {
                tokens: window.app.tokens,
                exchanges: window.app.exchanges
            }
        },
        methods: {
            getBalance: function(token) {
                if (token === "ETH") {
                    return "10.5";
                }
                return "$0.02";
            },
            getMarketBalance: (eid, token) => {
                return "yep!";
            }
        }
    }
</script>

<style scoped lang="scss">
    $red: #dd2266;

    .columns {
        margin: 0 auto;
        font-family: sans-serif;
        background-color: $red;
        padding: 10px;
        border-radius: 5px;
    }

    .cell {
        border: 1px solid $red;
        background-color: black;
        height: 60px;
        color: #eeeeee;
    }

    .cell-header {
        height: 80px;
    }

</style>
