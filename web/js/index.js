if (typeof window.app === "undefined" || window.app === null) window.app = {
    tokens: [],
    exchanges: []
}

import Vue from 'vue';
import Market from '../components/market';
import '../css/styles.scss';

let Vmarket = Vue.extend(Market);
new Vmarket().$mount("#market");
