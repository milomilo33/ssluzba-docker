# from django.shortcuts import render

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status
from rest_framework import permissions
from .models import Student, Profesor
from .serializers import StudentSerializer, ProfesorSerializer
from rest_framework.parsers import MultiPartParser, FormParser

import requests

class StudentApiView(APIView):
    parser_classes = (MultiPartParser, FormParser)

    def post(self, request, *args, **kwargs):

        data = {
            'jmbg': request.data.get('jmbg'), 
            'ime': request.data.get('ime'), 
            'prezime': request.data.get('prezime'), 
            'image_url': request.data.get('image_url')
        }
        serializer = StudentSerializer(data=data)

        if serializer.is_valid():
            # poslati zahtev uns aplikaciji
            r = requests.post('http://localhost:8003/api/uns/student', json={'jmbg': data["jmbg"], 'ime': data["ime"], 'prezime': data["prezime"]})
            if r.status_code != 201:
                print("UNS returned error.")
                return Response(status=status.HTTP_400_BAD_REQUEST)

            serializer.save()
            return Response(serializer.data, status=status.HTTP_201_CREATED)

        return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)


class ProfesorApiView(APIView):
    parser_classes = (MultiPartParser, FormParser)

    def post(self, request, *args, **kwargs):
        data = {
            'jmbg': request.data.get('jmbg'), 
            'ime': request.data.get('ime'), 
            'prezime': request.data.get('prezime'), 
            'image_url': request.data.get('image_url')
        }
        serializer = ProfesorSerializer(data=data)
        if serializer.is_valid():
            # poslati zahtev uns aplikaciji

            serializer.save()
            return Response(serializer.data, status=status.HTTP_201_CREATED)

        return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)