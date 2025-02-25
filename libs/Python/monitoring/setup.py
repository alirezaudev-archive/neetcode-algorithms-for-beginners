from setuptools import setup, find_packages

setup(
    name='monitoring',
    version='1.0',
    packages=find_packages(),
    install_requires=['psutil']
)