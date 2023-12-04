import request from '@/utils/request'
//获取基本信息
export function getInfo(params) {
    return request({
        url: '/base/info',
        method: 'get',
        params
    })
}

//获取基本信息
export function AddCategory(data) {
    return request({
        url: '/category/add_category',
        method: 'post',
        data
    })
}

//获取基本信息
export function UpdateCategory(data) {
    return request({
        url: '/category/update_category',
        method: 'post',
        data
    })
}

//获取基本信息
export function DeleteCategory(params) {
    return request({
        url: '/category/delete_category',
        method: 'get',
        params
    })
}

//获取基本信息
export function AddFileInfo(data) {
    return request({
        url: '/file_info/add_fileinfo',
        method: 'post',
        data
    })
}
//获取基本信息
export function UpdateFileInfo(data) {
    return request({
        url: '/file_info/update_fileinfo',
        method: 'post',
        data
    })
}

//获取基本信息
export function DeleteFileInfo(params) {
    return request({
        url: '/file_info/delete_fileinfo',
        method: 'get',
        params
    })
}


