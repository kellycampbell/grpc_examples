from __future__ import absolute_import, print_function

import grpc
import time

from concurrent import futures

from protos import catpic_pb2
from protos import catpic_pb2_grpc

_ONE_DAY_IN_SECONDS = 60 * 60 * 24


class CatPicServer(object):

    def __init__(self, catpic_service, server_port):
        self.server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))

        catpic_pb2_grpc.add_CatPicsServiceServicer_to_server(catpic_service, self.server)
        self.server.add_insecure_port('[::]:{server_port}'.format(server_port=server_port))

    def start(self):
        self.server.start()

    def stop(self):
        self.server.stop(0)

    def await_termination(self):
        """
        server.start() doesn't block so we explicitly block here unless someone keyboard-exits us.
        :return:
        """
        try:
            while True:
                time.sleep(_ONE_DAY_IN_SECONDS)
        except KeyboardInterrupt:
            self.server.stop(0)
        pass


class CatPicService(catpic_pb2_grpc.CatPicsServiceServicer):

    def GetCatPic(self, request, context):
        print("CatPic server received: " + request.cat_pic_id)
        print("  request = %r" % request)
        print("  auth_context = %r" % (context.auth_context(),))
        print("  metadata = %r" % (context.invocation_metadata(),))
        reply = catpic_pb2.CatPic()
        reply.url = 'hello'
        return reply


def main():
    catpic_server = CatPicServer(CatPicService(), 50051)
    catpic_server.start()
    catpic_server.await_termination()

if __name__ == '__main__':
    main()