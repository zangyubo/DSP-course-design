import axios from 'axios'

// 定义接收数据的接口
export interface Receive {
  message: string // 返回的消息内容
}

// 提交数据并接收响应（文件）
export function sendGetRequest(): Promise<void> {
  return axios
    .get('http://localhost:8080/testGetData', {
      responseType: 'blob', // 获取二进制数据（文件）
    })
    .then((response) => {
      // 获取返回的文件 Blob 对象
      const fileBlob = response.data

      // 创建一个临时的 URL 用于下载文件
      const fileUrl = URL.createObjectURL(fileBlob)

      // 创建一个临时的下载链接并触发下载
      const link = document.createElement('a')
      link.href = fileUrl
      link.download = 'time.npy' // 设置下载文件的名称
      link.click()

      // 释放临时的 URL 对象
      URL.revokeObjectURL(fileUrl)
    })
    .catch((error) => {
      console.error('Error fetching file:', error)
    })
}
