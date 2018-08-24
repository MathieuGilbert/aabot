if (typeof window.app === "undefined" || window.app === null) window.app = {
    tokens: [],
    exchanges: [],
    wallets: []
}

import Vue from 'vue';
import BigNumber from 'bignumber.js'
import Market from '../components/market-monitor';
import '../css/styles.scss';

window.BigNumber = BigNumber;
window.Vue = Vue;

let Vmarket = Vue.extend(Market);
new Vmarket().$mount("#market-monitor");
