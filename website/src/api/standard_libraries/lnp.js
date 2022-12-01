import request from '@/utils/request';
export function GetStdLnpLibraries(query) {
    return request({
        url: '/std/lnp',
        method: 'GET',
        params: query
    });
}
export function CreateStdLnpLibraries(lnp) {
    return request({
        url: '/std/lnp',
        method: 'POST',
        data: lnp
    });
}
export function UpdateStdLnpLibraries(lnp) {
    return request({
        url: '/std/lnp/' + lnp.id,
        method: 'PATCH',
        data: lnp
    });
}
export function DeleteStdLnpLibraries(id) {
    return request({
        url: '/std/lnp/' + id,
        method: 'DELETE'
    });
}
//# sourceMappingURL=lnp.js.map