import numpy as np
import matplotlib.pyplot as plt
from pydub import AudioSegment
from scipy.fft import fft, fftfreq

# 加载音频文件
audio = AudioSegment.from_mp3("../dist/test.mp3")

# 设置采样率为 22,050 Hz
audio = audio.set_frame_rate(22050)

# 获取音频样本数据
samples = np.array(audio.get_array_of_samples())
sample_rate = audio.frame_rate

# 将样本数据转换为 float32 类型，并保留三位小数
samples = np.round(samples.astype(np.float32), 3)


# 获取时域数据
def plot_time_domain(_samples, _sample_rate):
    # times = np.arange(len(_samples)) / _sample_rate
    # # 将 times 和 samples 数据转换为 float32 类型，并保留三位小数
    # times = np.round(times.astype(np.float32), 3)

    times = [1, 2, 3, 4, 5]
    

    

    # 保存时域数据为 .npy 文件
    np.save("../dist/time.npy", times)
    np.save("../dist/samples.npy", _samples)

    # 如果需要打印查看
    # print("时域数据（时间）:", times[:10])  # 打印前10个时间点
    # print("时域数据（样本）:", _samples[:10])  # 打印前10个样本值

    # 绘制时域图（可以取消注释）
    # plt.figure(figsize=(10, 6))
    # plt.plot(times, samples)
    # plt.title("Time Domain")
    # plt.xlabel("Time (s)")
    # plt.ylabel("Amplitude")
    # plt.grid(True)
    # plt.show()
#
#
# # 获取频域数据
# def plot_frequency_domain(samples, sample_rate):
#     N = len(samples)
#     T = 1.0 / sample_rate
#     yf = fft(samples)
#     xf = fftfreq(N, T)[:N // 2]
#
#     # 将频率和幅度数据转换为 float32 类型，并保留三位小数
#     xf = np.round(xf.astype(np.float32), 3)
#     amplitude = np.round(2.0 / N * np.abs(yf[:N // 2]).astype(np.float32), 3)
#
#     # 保存频域数据为 .npy 文件
#     np.save("../dist/frequency.npy", xf)
#     np.save("../dist/amplitude.npy", amplitude)
#
#     # 如果需要打印查看
#     print("频域数据（频率）:", xf[:10])  # 打印前10个频率值
#     print("频域数据（幅度）:", amplitude[:10])  # 打印前10个幅度值
#
#     # 绘制频域图（可以取消注释）
#     # plt.figure(figsize=(10, 6))
#     # plt.plot(xf, amplitude)
#     # plt.title("Frequency Domain")
#     # plt.xlabel("Frequency (Hz)")
#     # plt.ylabel("Amplitude")
#     # plt.grid(True)
#     # plt.show()
#
#
# # 绘制图形并保存数据
plot_time_domain(samples, sample_rate)
# plot_frequency_domain(samples, sample_rate)


# import numpy as np
# import os
# import matplotlib.pyplot as plt
# from pydub import AudioSegment
# from scipy.fft import fft, fftfreq
#
# # 文件路径
# time_file = "../dist/time.npy"
# samples_file = "../dist/samples.npy"
# frequency_file = "../dist/frequency.npy"
# amplitude_file = "../dist/amplitude.npy"
#
#
# # 检查文件是否存在，如果存在则删除
# def remove_if_exists(file_path):
#     if os.path.exists(file_path):
#         print(f"文件 {file_path} 已存在，正在删除...")
#         os.remove(file_path)
#
#
# # 加载音频文件
# audio = AudioSegment.from_mp3("../dist/test.mp3")
#
# # 设置采样率为 22,050 Hz
# audio = audio.set_frame_rate(22050)
#
# # 获取音频样本数据
# samples = np.array(audio.get_array_of_samples())
# sample_rate = audio.frame_rate
#
# # 将样本数据转换为 float16 类型，并保留三位小数
# samples = np.round(samples.astype(np.float16), 3)
#
#
# # 获取时域数据
# def plot_time_domain(samples, sample_rate):
#     times = np.arange(len(samples)) / sample_rate
#     # 将 times 转换为 float16 类型，并保留三位小数
#     times = np.round(times.astype(np.float16), 3)
#
#     # 删除已存在的文件（如果有的话）
#     remove_if_exists(time_file)
#     remove_if_exists(samples_file)
#
#     # 保存时域数据为 .npy 文件
#     np.save(time_file, times)
#     np.save(samples_file, samples)
#
#     # 如果需要打印查看
#     print("时域数据（时间）:", times[:10])  # 打印前10个时间点
#     print("时域数据（样本）:", samples[:10])  # 打印前10个样本值
#
#     # 绘制时域图（可以取消注释）
#     # plt.figure(figsize=(10, 6))
#     # plt.plot(times, samples)
#     # plt.title("Time Domain")
#     # plt.xlabel("Time (s)")
#     # plt.ylabel("Amplitude")
#     # plt.grid(True)
#     # plt.show()
#
#
# # 获取频域数据
# def plot_frequency_domain(samples, sample_rate):
#     N = len(samples)
#     T = 1.0 / sample_rate
#     yf = fft(samples)
#     xf = fftfreq(N, T)[:N // 2]
#
#     # 将频率和幅度数据转换为 float16 类型，并保留三位小数
#     xf = np.round(xf.astype(np.float16), 3)
#     amplitude = np.round(2.0 / N * np.abs(yf[:N // 2]).astype(np.float16), 3)
#
#     # 删除已存在的文件（如果有的话）
#     remove_if_exists(frequency_file)
#     remove_if_exists(amplitude_file)
#
#     # 保存频域数据为 .npy 文件
#     np.save(frequency_file, xf)
#     np.save(amplitude_file, amplitude)
#
#     # 如果需要打印查看
#     print("频域数据（频率）:", xf[:10])  # 打印前10个频率值
#     print("频域数据（幅度）:", amplitude[:10])  # 打印前10个幅度值
#
#     # 绘制频域图（可以取消注释）
#     # plt.figure(figsize=(10, 6))
#     # plt.plot(xf, amplitude)
#     # plt.title("Frequency Domain")
#     # plt.xlabel("Frequency (Hz)")
#     # plt.ylabel("Amplitude")
#     # plt.grid(True)
#     # plt.show()
#
#
# # 绘制图形并保存数据
# plot_time_domain(samples, sample_rate)
# plot_frequency_domain(samples, sample_rate)
