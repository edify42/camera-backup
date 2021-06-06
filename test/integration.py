import sys, hashlib, argparse

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

if __name__ == '__main__':
  # This makes me feel like python is a shitty language to write code with.
  parser = argparse.ArgumentParser(prog="Integration test suite")
  subparsers = parser.add_subparsers(dest='subparser')

  parser_check_file = subparsers.add_parser('calculate', help="Calculate the MD5 and e-tag values of a particular file")
  parser_check_file.add_argument('-f', '--file', dest='file', help='File we wish to check')
  
  parser_start_test = subparsers.add_parser('generate', help="Generate a bunch of test files in a certain directory")
  parser_start_test.add_argument('-d', '--destination', dest='destination', help='Destination directory of the new files')

  parser_create_mixed_files = subparsers.add_parser('middle', help="Middle test step which creates a mix of new and old files")
  parser_create_mixed_files.add_argument(
    '-s', '--source', dest='source', help='Source directory of old files')
  parser_create_mixed_files.add_argument(
    '-d', '--destination', dest='destination', help='Destination directory of mixed files')

  parser.add_argument("--log_level", help="Log level of the command line output", dest="log_level")

  kwargs = vars(parser.parse_args())
  
  if len(sys.argv) < 2 or len(sys.argv) > 5:
    parser.print_help()
    exit()

  if kwargs['subparser'] == 'calculate':
    source_path = kwargs['file']
    chunk_size = 8 # sys.argv[2] # Yeah we don't need to be clever here.
    chunk_size = int(chunk_size) * 1024 * 1024
    print(calculate_s3_etag(source_path,chunk_size))
