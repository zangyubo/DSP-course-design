from src.tempTest import tempTest
from src.enviromentUrl import get_environment

environment_url  = get_environment()
project_url = '/dspOfPython'

if __name__ == '__main__':
    _url = environment_url
    print(_url)

    tempTest(_url)
