import { defineMock } from '../_base'
import { getDelayTime, resultSuccess } from '../_utils'
import fileList from '../_data/file_list'

const getList = (type: number | string) => {
  if (Number(type) === 0) return fileList
  const res: any[] = []
  const FileMap = {
    1: ['jpg', 'png', 'jpeg', 'gif'],
    2: ['txt', 'doc', 'xls'],
    3: ['mp4'],
    4: ['mp3'],
    5: ['zip', 'rar', 'ppt', 'css', 'js', 'html']
  }
  const arr = FileMap[type as keyof typeof FileMap] || []
  arr.forEach((item) => {
    const data = fileList.filter((i) => i.extendName === item)
    res.push(...data)
  })
  return res
}

export default defineMock([
  {
    url: '/file/list',
    method: 'get',
    timeout: getDelayTime(),
    response: ({ query }) => {
      const { fileType } = query
      const list = getList(fileType)
      return resultSuccess({
        total: list.length,
        records: JSON.parse(JSON.stringify(list))
      })
    }
  }
])
