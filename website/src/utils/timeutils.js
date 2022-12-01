import dayjs from 'dayjs';
export function UnixTime2HumanWithUnix(timestamp) {
    return dayjs.unix(timestamp).format('YYYY-MM-DD HH:mm:ss');
}
export function UnixTime2HumanWithStr(str) {
    // UnixTime2Human(timestamp)
    return dayjs(str).format('YYYY-MM-DD HH:mm:ss');
}
export default UnixTime2HumanWithStr;
//# sourceMappingURL=timeutils.js.map