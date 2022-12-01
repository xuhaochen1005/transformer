import request from '@/utils/request';
export function GetStdlcdLibraries(query) {
    return request({
        url: '/std/lcd/',
        method: 'GET',
        params: query
    });
}
export function DeleteStdlcdLibraries(id) {
    return request({
        url: '/std/lcd/' + id,
        method: 'DELETE'
    });
}
export function UpdateStdlcdLibraries(lcd) {
    return request({
        url: '/std/lcd/' + lcd.id,
        method: 'PATCH',
        data: lcd
    });
}
export function CreateStdlcdLibraries(lcd) {
    return request({
        url: '/std/lcd',
        method: 'POST',
        data: lcd
    });
}
//# sourceMappingURL=lcd.js.map