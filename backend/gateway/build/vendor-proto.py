import io
import os.path
import re
import shutil
import sys

VENDOR_PROTO_DIR = 'vendor.proto'

service_name_regex = r"github\.com\/knowdateapp\/knowdateapp\/backend(?P<service_name>\/\w+\/)"
service_name_pattern = re.compile(service_name_regex)

protobuf_target_dir_regex = r"github\.com\/knowdateapp\/knowdateapp\/backend\/.*(?P<pb_dir>\/internal\/pb\/)"
protobuf_target_dir_pattern = re.compile(protobuf_target_dir_regex)

for dependency in sys.stdin:
    source_file = dependency.strip()
    source_filename = os.path.basename(source_file)

    target_directory = os.path.join('.', VENDOR_PROTO_DIR, os.path.dirname(source_file.strip('../')))
    target_file = os.path.join(target_directory, source_filename)

    print(f'Copy file {source_file} to {target_directory} directory.')

    os.makedirs(target_directory, exist_ok=True)
    shutil.copy(source_file, target_directory)

    content = io.StringIO()
    with open(target_file, 'r') as file:
        for line in file:
            service_name_match = service_name_pattern.search(line)
            if not service_name_match:
                content.write(line)
                continue
            protobuf_target_dir_match = protobuf_target_dir_pattern.search(line)
            if not protobuf_target_dir_match:
                continue
            service_name = service_name_match.group('service_name')
            pb_dir = protobuf_target_dir_match.group('pb_dir')
            line = line.replace(service_name, '/gateway/')
            line = line.replace(pb_dir, f'/{pb_dir.strip("/")}/{service_name.strip("/")}/')
            content.write(line)

    with open(target_file, 'w') as file:
        content.seek(0)
        shutil.copyfileobj(content, file)
