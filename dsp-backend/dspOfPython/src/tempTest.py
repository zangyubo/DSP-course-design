import numpy as np

def tempTest(store_url):
    times = [1,2,3,4,5]

    times_file_url = store_url + "/dist/time.npy"

    print(times_file_url)
    np.save(times_file_url, times)
