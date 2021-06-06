import os, hashlib, argparse
from typing import Union

def calculate_s3_etag(file_path, chunk_size=8 * 1024 * 1024):
  md5s = []

  with open(file_path, 'rb') as fp:
    while True:
      data = fp.read(chunk_size)
      if not data:
        break
      md5s.append(hashlib.md5(data))

  if len(md5s) < 1:
    return '{}'.format(hashlib.md5().hexdigest())

  if len(md5s) == 1:
    return '{}'.format(md5s[0].hexdigest())

  digests = b''.join(m.digest() for m in md5s)
  digests_md5 = hashlib.md5(digests)
  return '{}-{}'.format(digests_md5.hexdigest(), len(md5s))

def calculate_md5(data: Union[str, bytes]):
  return hashlib.md5(data).hexdigest()

def calculate_checksum(file: str, md5 = True):
  if not os.path.exists(file):
    raise ValueError('{} is non-existant'.format(file))
  if md5:
    with open(file, 'rb') as fp:
      data = fp.read()
      return calculate_md5(data)
  else:
    return calculate_s3_etag(file)

    

def generate(dest_path: str, md5 = False, num_files = 10, size_multiplier = 16):
  base_size = 1024 * 1024 # 1MB
  for i in range(num_files):
    file = 'temp_file_{}'.format(i)
    qualified_file = os.path.join(dest_path, file)
    with open(qualified_file, 'wb') as fout:
      fout.write(os.urandom(base_size * size_multiplier)) 
    checksum = calculate_checksum(qualified_file, md5)
    new_file = os.path.join(dest_path, checksum)
    os.rename(qualified_file, new_file)


if __name__ == '__main__':
  # This makes me feel like python is a shitty language to write code with.
  parser = argparse.ArgumentParser(prog="Integration test suite")
  subparsers = parser.add_subparsers(dest='subparser')

  parser_check_file = subparsers.add_parser('calculate', help="Calculate the MD5 and e-tag values of a particular file")
  parser_check_file.add_argument('-f', '--file', dest='file',
    help='File we wish to check', required=True)
  parser_check_file.add_argument('-m', '--md5',
    action=argparse.BooleanOptionalAction,
    help="Return md5 sum instead of e-tag")
  
  parser_start_test = subparsers.add_parser('generate', help="Generate a bunch of test files in a certain directory")
  parser_start_test.add_argument(
    '-d',
    '--destination',
    dest='destination',
    help='Destination directory of the new files',
    required=True)
  parser_start_test.add_argument('-c', '--count', dest='count',
    type=int,
    help='Number of files to generate')
  parser_start_test.add_argument('-m', '--md5',
    action=argparse.BooleanOptionalAction,
    help="Generate md5 sum instead of e-tag")

  parser_create_mixed_files = subparsers.add_parser('middle', help="Middle test step which creates a mix of new and old files")
  parser_create_mixed_files.add_argument(
    '-s', '--source', dest='source', help='Source directory of old files')
  parser_create_mixed_files.add_argument(
    '-d', '--destination', dest='destination', help='Destination directory of mixed files')

  # haven't figured out how to do global args - shitty python...
  # parser.add_argument("--log_level", help="Log level of the command line output", dest="log_level")

  kwargs = vars(parser.parse_args())
  
  if not kwargs['subparser']:
    parser.print_help()
    exit()

  if kwargs['subparser'] == 'calculate':
    source_path = kwargs['file']
    if kwargs['md5']:
      with open(source_path, 'rb') as fp:
        data = fp.read()
        print(calculate_md5(data))

    else:
      chunk_size = 8 # Yeah we don't need to be clever here.
      chunk_size = int(chunk_size) * 1024 * 1024
      print(calculate_s3_etag(source_path, chunk_size))

  if kwargs['subparser'] == 'generate':
    dest_path = kwargs['destination']
    file_count = 10
    if kwargs['count']:
      file_count = kwargs['count']
    md5 = kwargs['md5']
    if os.path.isdir(dest_path):
      print('generating fake data to {}'.format(dest_path))
      generate(dest_path, md5, file_count)
    else:
      print('no such directory {} - please create it first'.format(dest_path))
      exit(1)