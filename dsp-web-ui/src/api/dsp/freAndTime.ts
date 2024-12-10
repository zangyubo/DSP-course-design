import axios from 'axios';
import npyjs from 'npyjs';

const Npyjs = npyjs;

type channel = 'freX' | 'freY' | 'timeX' | 'timeY' | 'picture';

export interface UploadFileData {
  file: File; // send file
  channel: channel;
}

export interface UploadFileRes {
  message: string;
}

export function uploadFile(data: UploadFileData): Promise<UploadFileRes> {
  const formData = new FormData();

  formData.append('file', data.file);
  formData.append('channel', data.channel);

  return axios
    .post<UploadFileRes>(
      'http://localhost:8080/freAndTimeChartUpload',
      formData,
      {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      }
    )
    .then((res) => res.data);
}

export interface receiveFileData {
  channel: channel;
}

export interface receiveFileRes {
  message: string;
  fileData?: {
    dtype: 'float32';
    shape: number[];
    data: Float32Array;
    fortranOrder: boolean;
  };
}

export function receiveFile(data: receiveFileData): Promise<receiveFileRes> {
  return axios
    .get(`http://localhost:8080/basicChartReceiveMP3${data}`, {
      responseType: 'blob',
    })
    .then((res) => {
      const fileBlob = res.data;

      const reader = new FileReader();

      return new Promise((resolve) => {
        reader.onload = () => {
          const npyLoader = new Npyjs();

          const arrayBuffer = reader.result as ArrayBuffer;

          npyLoader
            .load(arrayBuffer)
            .then((returnData) => {
              resolve({
                message: 'success',
                fileData: returnData as receiveFileRes['fileData'],
              });
            })
            .catch((err) => {
              // eslint-disable-next-line no-console
              console.log(err);
              resolve({ message: 'error' });
            });
        };

        reader.readAsArrayBuffer(fileBlob);
      });
    });
}
