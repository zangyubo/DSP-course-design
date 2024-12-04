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

% 用户输入密钥
key = input('请输入密钥字符串: ', 's');

% 将密钥转换为 ASCII 值
key_ascii = double(key);

% 加密音频信号 (异或加密)
y_encrypted = y;
key_length = length(key_ascii);
for i = 1:length(y)
    % 对音频每个采样点进行加密
    y_encrypted(i) = bitxor(y(i), key_ascii(mod(i-1, key_length) + 1)); 
end

% 保存加密后的音频文件
encrypted_filename = 'encrypted_audio.wav';
audiowrite(encrypted_filename, y_encrypted, Fs);
disp(['加密音频已保存为: ', encrypted_filename]);

% 播放加密后的音频
disp('播放加密后的音频...');
sound(y_encrypted, Fs);
