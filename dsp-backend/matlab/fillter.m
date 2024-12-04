% 用户选择音频文件
[file, path] = uigetfile({'*.wav;*.mp3', '音频文件 (*.wav, *.mp3)'}, '请选择一个音频文件');
if isequal(file, 0)
    disp('用户取消了选择');
    return;
end
filename = fullfile(path, file);

% 读取音频文件
[y, Fs] = audioread(filename); % y: 音频信号, Fs: 采样率

% 如果音频是多通道，选择一个通道
if size(y, 2) > 1
    y = y(:, 1);
end

% 获取时间轴
t = (0:length(y)-1) / Fs;

% 显示模式选择菜单
mode = input('请输入模式 (1: IIR 滤波, 2: FIR 滤波): ');

% 滤波器设计参数
Fc = 1000; % 截止频率 (Hz)
N = 50;    % 滤波器阶数

% 根据模式进行滤波
switch mode
    case 1 % IIR 滤波
        % 使用巴特沃斯低通滤波器
        [b, a] = butter(4, Fc / (Fs / 2)); % 归一化频率
        y_filtered = filter(b, a, y);

    case 2 % FIR 滤波
        % 使用窗函数设计的FIR低通滤波器
        b = fir1(N, Fc / (Fs / 2)); % 归一化频率
        y_filtered = filter(b, 1, y);

    otherwise
        disp('无效的模式');
        return;
end

% 绘制原始和滤波后信号
figure;
subplot(2, 1, 1);
plot(t, y);
title('原始音频信号');
xlabel('时间 (s)');
ylabel('振幅');
grid on;

subplot(2, 1, 2);
plot(t, y_filtered);
title(['滤波后音频信号 (' (mode == 1) * "IIR" + (mode == 2) * "FIR" ')']);
xlabel('时间 (s)');
ylabel('振幅');
grid on;

% 播放原始和滤波后音频
disp('播放原始音频...');
sound(y, Fs);
pause(length(y) / Fs + 1); % 等待播放完成

disp('播放滤波后音频...');
sound(y_filtered, Fs);

% 绘制频谱对比
figure;
NFFT = 2^nextpow2(length(y)); % 确保频率分辨率足够
f = Fs * (0:(NFFT/2)) / NFFT; % 频率轴

% 计算原始信号频谱
Y = fft(y, NFFT) / length(y);
P1 = abs(Y(1:NFFT/2+1));

% 计算滤波后信号频谱
Y_filtered = fft(y_filtered, NFFT) / length(y_filtered);
P1_filtered = abs(Y_filtered(1:NFFT/2+1));

% 绘制频谱
plot(f, 20*log10(P1), 'b', 'LineWidth', 1.2); hold on;
plot(f, 20*log10(P1_filtered), 'r', 'LineWidth', 1.2);
title('频谱对比');
xlabel('频率 (Hz)');
ylabel('幅值 (dB)');
legend('原始信号', '滤波后信号');
grid on;
