#coding=utf-8


from zipfile import ZipFile
from zipfile import ZIP_STORED
import os
import sys
from shutil import copyfileobj

curpath = os.path.dirname(os.path.join(__file__))


def package_sfx(target_filepath, sfx_filepath, dir0, entry):
    '''
    :param target_filepath: The final self extracting file path, Windows exe
    :param sfx_filepath:  The prepare SFX file, Windows exe
    :param dir0: The directory of files to be packaged
    :param entry: The final self extracting file exec command
    :return:
    '''

# Path=.\temp\r
# Overwrite=1\r
#
    sfx_config = ''';The comment below contains SFX script commands\r
Setup={}\r
TempMode\r
Silent=1\r
'''.format(entry)

    files = []
    for root,sub, fs in os.walk(dir0):
        if not root.startswith(dir0):
            raise ValueError("{} not startswith {}".format(root, dir0))
        sub_root = root[len(dir0)::]
        for f in fs:
            # (absolute path, filename in zip)
            files.append((os.path.join(root,f),
                          os.path.join(sub_root,f)))

    with open(target_filepath, "wb") as fw:
        with open(sfx_filepath, "rb") as fr:
            copyfileobj(fr, fw)

    with ZipFile(target_filepath, "a", compression=ZIP_STORED) as fw:
        for f in files:
            fw.write(filename=f[0], arcname=f[1])
        fw.comment = sfx_config


def main():
   pass


if __name__ == '__main__':
    main()

