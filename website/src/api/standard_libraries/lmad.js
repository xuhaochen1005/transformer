import request from '@/utils/request';
export function GetStdLmadLibraries(query) {
    return request({
        url: '/std/lmad',
        method: 'GET',
        params: query
    });
}
export function CreateStdLmadLibraries(lmad) {
    return request({
        url: '/std/lmad',
        method: 'POST',
        data: lmad
    });
}
export function UpdateStdLmadLibraries(lmad) {
    return request({
        url: '/std/lmad/' + lmad.id,
        method: 'PATCH',
        data: lmad
    });
}
export function DeleteStdLmadLibraries(id) {
    return request({
        url: '/std/lmad/' + id,
        method: 'DELETE'
    });
}
//# sourceMappingURL=lmad.js.map