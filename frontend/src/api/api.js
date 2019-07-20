import axios from 'axios'
export const test = 'http://192.168.248.188';

// 登录
export const requestLogin = params => { return axios.post('v1/user/login', params).then(res => res.data) }