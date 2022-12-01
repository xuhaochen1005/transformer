import request from '@/utils/request';
export function GetStdLshvLibraries(query) {
    return request({
        url: '/std/lshv/',
        method: 'GET',
        params: query
    });
}
export function DeleteStdLshvLibraries(id) {
    return request({
        url: '/std/lshv/' + id,
        method: 'DELETE'
    });
}
export function UpdateStdLshvLibraries(lshv) {
    return request({
        url: '/std/lshv/' + lshv.id,
        method: 'PATCH',
        data: lshv
    });
}
export function CreateStdLshvLibraries(lshv) {
    return request({
        url: '/std/lshv',
        method: 'POST',
        data: lshv
    });
}
//# sourceMappingURL=lshv.js.map