// var fly=require("flyio")
// import {Toast,Dialog,Notify} from 'vant';
// import 'vant/lib/index.css';

fly.config.baseURL = location.protocol + '//' + location.host
fly.config.timeout=10000;

fly.config.headers = {
    'Content-Type': 'application/x-www-form-urlencoded'
}

//添加请求拦截器
fly.interceptors.request.use((request)=>{
    //给所有请求添加自定义header
    // request.headers["X-Tag"]="flyio";
    //打印出请求体
    // console.log('request',request.body)
    // 终止请求
    //var err=new Error("xxx")
    //err.request=request
    //return Promise.reject(new Error(""))

    //可以显式返回request, 也可以不返回，没有返回值时拦截器中默认返回request
    return request;
})

//添加响应拦截器，响应拦截器会在then/catch处理之前执行
fly.interceptors.response.use(
    (response) => {
        //只将请求结果的data字段返回
        return response.data
    },
    (err) => {
        //发生网络错误后会走到这里
        //return Promise.resolve("ssss")
        // Toast.clear()
        // Notify({ type: 'danger', message: '网络错误'});
    }
)

// /**
//  * get方法，对应get请求
//  * @param {String} url [请求的url地址]
//  * @param {Object} params [请求时携带的参数]
//  */
// export function get(url, params){
//     return new Promise((resolve, reject) =>{
//         fly.get(
//             url,
//             params
//         ).then(res => {
//             resolve(res.data);
//         }).catch(err =>{
//             reject(err.data)
//         })
//     });
// }

// /**
//  * post方法，对应post请求
//  * @param {String} url [请求的url地址]
//  * @param {Object} params [请求时携带的参数]
//  */
// export function post(url, params) {
//     return new Promise((resolve, reject) => {
//         fly.post(url, QS.stringify(params))
//         .then(res => {
//             resolve(res.data);
//         })
//         .catch(err =>{
//             reject(err.data)
//         })
//     });
// }

// export default fly