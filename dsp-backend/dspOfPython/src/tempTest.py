import numpy as np

def tempTest(store_url):
    times = [1,2,3,4,5]
    value = [5,4,3,2,1]

    times_file_url = store_url + "/dist/time.npy"
    value_file_url = store_url + "/dist/value.npy"
    np.save(times_file_url, times)
    np.save(value_file_url, value)
