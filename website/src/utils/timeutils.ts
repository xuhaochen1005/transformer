
import dayjs from 'dayjs'

export function UnixTime2HumanWithUnix(timestamp:number): string {
  return dayjs.unix(timestamp).format('YYYY-MM-DD HH:mm:ss')
}
export function UnixTime2HumanWithStr(str: string) {
  // UnixTime2Human(timestamp)
  return dayjs(str).format('YYYY-MM-DD HH:mm:ss')
}

export default UnixTime2HumanWithStr
