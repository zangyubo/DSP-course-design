import numpy as np
import argparse
from pydub import AudioSegment
from scipy.fft import fft, fftfreq

def basicChart(channel: int, mode: int):
    # 加载音频文件
    if channel == 1:
        audio = AudioSegment.from_mp3("E:/DSP-course-design/dsp-backend/dist/basicChart/channel1/signal.mp3")
    else:
        audio = AudioSegment.from_mp3("E:/DSP-course-design/dsp-backend/dist/basicChart/channel2/signal.mp3")

    # 设置采样率为 10000 Hz
    audio = audio.set_frame_rate(10000)

    # 获取音频的样本数据并转为 NumPy 数组
    samples = np.array(audio.get_array_of_samples())
    sample_rate = audio.frame_rate

    # 转换样本数据为 float32 类型
    samples = np.round(samples.astype(np.float32), 3)

    # 获取时域数据
    time_axis = np.linspace(0, len(samples) / sample_rate, num=len(samples))

    # 获取频域数据
    N = len(samples)
    frequency_axis = fftfreq(N, 1 / sample_rate)
    frequency_data = fft(samples)

    # 只保留正频率部分
    positive_frequencies = frequency_axis[:N//2]
    positive_amplitudes = np.abs(frequency_data)[:N//2]

    # 将时域和频域数据存储在字典中
    data_dict = {
        'time_axis': time_axis,
        'samples': samples,
        'positive_frequencies': positive_frequencies,
        'positive_amplitudes': positive_amplitudes
    }

    if mode == 1:
        save_path = f"E:/DSP-course-design/dsp-backend/dist/basicChart/channel{channel}/channel{channel}_freXdata.npy"
        np.save(save_path, np.round(positive_frequencies, 3).astype(np.float32))
    elif mode == 2:
        save_path = f"E:/DSP-course-design/dsp-backend/dist/basicChart/channel{channel}/channel{channel}_freYdata.npy"
        np.save(save_path, np.round(positive_amplitudes, 3).astype(np.float32))
    elif mode == 3:
        save_path = f"E:/DSP-course-design/dsp-backend/dist/basicChart/channel{channel}/channel{channel}_TimeXdata.npy"
        np.save(save_path, np.round(time_axis, 3).astype(np.float32))
    elif mode == 4:
        save_path = f"E:/DSP-course-design/dsp-backend/dist/basicChart/channel{channel}/channel{channel}_TimeYdata.npy"
        np.save(save_path, np.round(samples, 3).astype(np.float32))

    print(f"数据已保存")

def main():
    # 设置命令行参数解析
    parser = argparse.ArgumentParser(description="Process audio and generate charts for specified channel and mode.")
    parser.add_argument('channel', type=int, choices=[1, 2], help='Specify the channel number (1 or 2)')
    parser.add_argument('mode', type=int, choices=[1, 2, 3, 4], help='Specify the mode (1: freX, 2: freY, 3: TimeX, 4: TimeY)')

    # 解析命令行参数
    args = parser.parse_args()

    # 调用函数处理数据
    basicChart(args.channel, args.mode)

if __name__ == '__main__':
    main()
