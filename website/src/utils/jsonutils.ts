export function stringify(obj: Object) : string {
  return JSON.stringify(obj, (k : string, v : any) : any => {
    // k,v分别代表传入对象的属性名和属性值
    // 判断k是因为replace会传入一个 传入参数对象 的值,但是没有k, 我们直接返回即可
    if (!k) {
      return v
    }
    if (typeof v === 'object') {
      return null
    }
    return v
  })
}

export function decodeBytesData(bufferStr : string) {
  try {
    return JSON.parse(Buffer.from(bufferStr, 'base64').toString())
  } catch (e) {
    console.log(e)
  }
}

export function encodeBytesData(data : any) : string {
  try {
    return Buffer.from(JSON.stringify(data)).toString('base64')
  } catch (e) {
    console.log(e)
  }
  return ''
}

export function decodeObjectBytes<T>(obj : T) {
  if (Array.isArray(obj)) {
    for (const item of obj) {
      decodeObjectBytes(item)
    }
    return
  }
  for (const f in obj) {
    if (typeof obj[f] === 'string') {
      try {
        obj[f] = JSON.parse(Buffer.from((obj[f] as unknown as string), 'base64').toString())
      } catch (e) {
      }
    }
    if (Array.isArray(obj[f])) {
      decodeObjectBytes(obj[f])
    }
  }
}
