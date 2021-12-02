from setuptools import setup, find_packages

setup(
    name='aoc',
    version='0.1',
    packages=find_packages(),
    include_package_data=True,
    install_requires=[
        'runit',
    ],
    entry_points='''
        [console_scripts]
        aoc=main:run
    ''',
)