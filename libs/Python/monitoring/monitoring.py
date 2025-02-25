from psutil import Process
from time import time
from typing import Callable
from os import getpid

def execute(func: Callable, *args, **kwargs):
    start, mem_before = time(), process_memory()
    result = func(*args, **kwargs)
    end, mem_after = time(), process_memory()

    print("execution time:", end - start)
    print("Memory usage  :", mem_after - mem_before, "MB")

    return result

def process_memory():
    process = Process(getpid())
    mem_info = process.memory_info()
    return mem_info.rss / (1024 ** 2)
