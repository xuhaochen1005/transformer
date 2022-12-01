import request from '@/utils/request';
export function GetStdLtfLibraries(query) {
    return request({
        url: '/std/ltf/',
        method: 'GET',
        params: query
    });
}
export function DeleteStdLtfLibraries(id) {
    return request({
        url: '/std/ltf/' + id,
        method: 'DELETE'
    });
}
export function UpdateStdLtfLibraries(ltf) {
    return request({
        url: '/std/ltf/' + ltf.id,
        method: 'PATCH',
        data: ltf
    });
}
export function CreateStdLtfLibraries(ltf) {
    return request({
        url: '/std/ltf',
        method: 'POST',
        data: ltf
    });
}
//# sourceMappingURL=ltf.js.map