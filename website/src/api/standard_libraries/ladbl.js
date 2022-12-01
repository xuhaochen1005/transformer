import request from '@/utils/request';
export function GetStdladblLibraries(query) {
    return request({
        url: '/std/ladbl/',
        method: 'GET',
        params: query
    });
}
export function DeleteStdladblLibraries(id) {
    return request({
        url: '/std/ladbl/' + id,
        method: 'DELETE'
    });
}
export function UpdateStdladblLibraries(ladbl) {
    return request({
        url: '/std/ladbl/' + ladbl.id,
        method: 'PATCH',
        data: ladbl
    });
}
export function CreateStdladblLibraries(ladbl) {
    return request({
        url: '/std/ladbl',
        method: 'POST',
        data: ladbl
    });
}
//# sourceMappingURL=ladbl.js.map