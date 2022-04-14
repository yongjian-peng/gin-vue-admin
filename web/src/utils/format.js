import { formatTimeToStr } from '@/utils/date'
import { getDict } from '@/utils/dictionary'

export const formatBoolean = (bool) => {
  if (bool !== null) {
    return bool ? '是' : '否'
  } else {
    return ''
  }
}
export const formatDate = (time) => {
  if (time !== null && time !== '') {
    var date = new Date(time)
    return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss')
  } else {
    return ''
  }
}

export const unixTime = (timeunix) => {
  const time = new Date(parseInt(timeunix) * 1000);
  const y = time.getFullYear();
  const m = time.getMonth() + 1;//getMonth获取的月份为（0--11），实际月份需要+1得出
  const d = time.getDate();
  const h = time.getHours();
  const mm = time.getMinutes();
  const s = time.getSeconds();
  const realtime = y+'-'+addZero(m)+'-'+addZero(d) +' ' +addZero(h) + ':' + addZero(mm) + ':' + addZero(s);
  return realtime;
}

/**
 * 添加0值
 */
export const addZero = (m) => {
  return m < 10 ? '0' + m : m;
}

export const filterDict = (value, options) => {
  const rowLabel = options && options.filter(item => item.value === value)
  return rowLabel && rowLabel[0] && rowLabel[0].label
}

export const getDictFunc = async(type) => {
  const dicts = await getDict(type)
  return dicts
}
