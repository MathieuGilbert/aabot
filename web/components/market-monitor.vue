<template>
    <div class="columns">
        <div class="column is-narrow">
            <div class="cell cell-header">
                &nbsp;
            </div>
            <div class="cell" v-for="token in tokens" :key="token.id">
                {{token.name}}
            </div>
        </div>
        <div class="column is-narrow">
            <div class="cell cell-header">
                <span class="icon is-small on" @click="updateWallets">
                    <i class="fab fa-ethereum"></i>
                </span>
                <span>Wallet</span>
            </div>
            <div class="cell" v-for="wallet in wallets" :key="wallet.id">
                <p>{{toEth(wallet.balance)}}</p>
            </div>
        </div>
        <div class="column is-narrow" v-for="exchange in exchanges" :key="exchange.id">
            <div class="cell cell-header">
                <span class="icon is-small on" @click="updateMarkets">
                    <i class="fas fa-sync"></i>
                </span>
                <span>{{exchange.name}}</span>
            </div>
            <div class="cell" v-for="market in exchange.markets" :key="market.id">
                <p>{{toEth(market.balance)}}</p>

                <div v-if="market.token.name != 'ETH'">
                    <span v-bind:class="[market.active ? 'on' : 'off', 'icon is-small']">
                        <i class="fas fa-power-off"></i>
                    </span>
                    <ul v-if="market.orders">
                        <li v-for="order in market.orders" :key="order.id" class="">
                            <span v-bind:class="[order.is_buy ? 'on' : 'off']">{{parseFloat(order.price).toFixed(6)}}</span>
                            <span class="">{{parseFloat(toEth(order.volume)).toFixed(3)}}</span>
                        </li>
                    </ul>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        name: "market-monitor",
        data() {
            return {
                tokens: window.app.tokens,
                exchanges: window.app.exchanges,
                wallets: window.app.wallets,
                socket: ''
            }
        },
        methods: {
            updateWallets: () => {
                $.ajax({
                    url: '/wallets',
                    type: 'GET',
                    cache: false,
                    contentType: false,
                    processData: false,
                    timeout: 5000
                }).done(function(data) {
                    window.app.wallets.splice(0, window.app.wallets.length);
                    for (var i = 0; i < data.wallets.length; i++) {
                        window.app.wallets.push(data.wallets[i]);
                    }
                }).fail(function(xhr, status, error) {
                    console.log("Error loading wallets");
                });
            },
            updateMarkets: () => {
                $.ajax({
                    url: '/markets',
                    type: 'GET',
                    cache: false,
                    contentType: false,
                    processData: false,
                    timeout: 5000
                }).done(function(data) {
                    for (var i = 0; i < window.app.exchanges.length; i++) {
                        var exchange = window.app.exchanges[i];
                        exchange.markets.splice(0, exchange.markets.length);

                        for (var i = 0; i < data.markets.length; i++) {
                            var market = data.markets[i];
                            if (market.exchange_id === exchange.id) {
                                exchange.markets.push(market);
                            }
                        }
                    }
                }).fail(function(xhr, status, error) {
                    console.log("Error loading markets");
                });
            },
            getOrderbooks: () => {
                for (var e = 0; e < window.app.exchanges.length; e++) {
                    let exchange = window.app.exchanges[e];

                    for (var m = 0; m < exchange.markets.length; m++) {
                        let mid = exchange.markets[m].id;

                        if (!exchange.markets[m].active) {
                            console.log(m + " is not active market. id:" + mid);
                            continue;
                        }

                        $.ajax({
                            url: '/orderbook?mid=' + mid + '&n=' + 5,
                            type: 'GET',
                            cache: false,
                            contentType: false,
                            processData: false,
                            timeout: 5000
                        }).done(function(data) {
                            console.log("setting orders for market: " + mid);

                            var e = window.app.exchanges.filter(e => {
                                return e.id === exchange.id;
                            })[0];

                            var m = e.markets.filter(m => {
                                return m.id = mid;
                            })[0];

                            if (m.orders === undefined || m.orders === null) {
                                console.log('~~~~~~~~~~~~')
                                m.orders = [];
                                window.Vue.set(m.orders, 'orders', data.orders)
                            }
                        }).fail(function(xhr, status, error) {
                            console.log("Error loading orderbook for market " + market);
                        });
                    }
                }
            },
            orderByTokenId: (list) => {
                list.sort((a, b) => {
                    if (a.token_id === undefined) {
                        return a.id > b.id ? -1 : 1;
                    }
                    return a.token_id > b.token_id ? -1 : 1;
                });
            },
            toEth: (wei) => {
                return new window.BigNumber(wei, 10).div(1000000000000000000).toString(10);
            },
            pingServer() {
                this.socket.send(JSON.stringify({'message': "wallet", "value": window.app.wallets[0]}));
            }
        },
        watch: {
            tokens: {
                handler: function(v, ov) {
                    this.updateWallets();
                    this.updateMarkets();
                },
                deep: true
            }
        },
        mounted() {
            this.socket = new WebSocket('ws://localhost:5000/ws');

            this.socket.addEventListener('message', function (event) {
                try {
                    var data = JSON.parse(event.data);

                    switch (data.message) {
                        case "wallet":
                            console.log("Got a wallet: " + JSON.stringify(data.value));
                            break;
                        default:
                            console.log('Unknown response: ' + JSON.stringify(data));
                    }
                } catch(e) {
                    console.log('Invalid JSON response: ' + event.data);
                }
            });
        }
    }
</script>

<style scoped lang="scss">
    $black: #000000;
    $slate: #202020;
    $green: #00ff00;
    $red: #ff0000;

    .columns {
        margin: 0 auto;
        font-family: sans-serif;
        background-color: $black;
        padding: 10px;
        border-radius: 5px;
    }

    .column {
        padding: 0px;
    }

    .cell {
        border: 1px solid $slate;
        background-color: black;
        height: 320px;
        padding: 10px;
        color: #eeeeee;
    }

    .cell-header {
        height: 80px;
    }

    span {
        &.on {
            color: $green;
        }

        &.off {
            color: $red;
        }

        &.icon:hover {
            background-color: $slate;
            cursor: pointer;
        }
    }
</style>
