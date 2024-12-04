% 读取 MP3 文件
[file, path] = uigetfile('*.mp3', '请选择一个MP3文件'); % 弹出对话框选择文件
if isequal(file, 0)
    disp('用户取消了选择');
    return;
end
filename = fullfile(path, file);

[y, Fs] = audioread(filename); % 读取音频信号和采样频率

% 如果音频是多通道，选择一个通道（例如第一个）
if size(y, 2) > 1
    y = y(:, 1);
end

% 获取时间轴
t = (0:length(y)-1) / Fs;

% 绘制时域波形
figure;
subplot(2, 1, 1); % 两行一列，当前为第一幅图
plot(t, y);
title('时域波形');
xlabel('时间 (s)');
ylabel('振幅');
grid on;

% 计算频域
N = length(y); % 数据点数
f = (0:N-1) * (Fs / N); % 频率轴
Y = fft(y); % 快速傅里叶变换
P2 = abs(Y / N); % 双边幅度谱
P1 = P2(1:floor(N/2)+1); % 单边幅度谱
P1(2:end-1) = 2 * P1(2:end-1); % 适当放大非直流分量

% 绘制频域频谱
subplot(2, 1, 2); % 当前为第二幅图
plot(f(1:floor(N/2)+1), P1);
title('频域频谱');
xlabel('频率 (Hz)');
ylabel('幅值');
grid on;
