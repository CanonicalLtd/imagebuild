// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

import axios from 'axios'
import constants from './constants'

let service = {
    boardsList:  (query, cancelCallback) => {
        return axios.get(constants.baseUrl + 'boards');
    },

    buildRequest:  (query, cancelCallback) => {
        return axios.post(constants.baseUrl + 'build', query);
    },

    buildFetch:  (query, cancelCallback) => {
        return axios.post(constants.baseUrl + 'build/fetch', query);
    },

    storeSearch: (snapName, bld, cancelCallback) => {
        return axios.post(constants.baseUrl + 'store/snaps/' + snapName, bld);
    },
}

export default service