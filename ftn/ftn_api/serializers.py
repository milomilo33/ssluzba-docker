from rest_framework import serializers
from .models import Student, Profesor

class StudentSerializer(serializers.ModelSerializer):
    image_url = serializers.ImageField(required=False)

    class Meta:
        model = Student
        fields = ["jmbg", "ime", "prezime", "image_url"]


class ProfesorSerializer(serializers.ModelSerializer):
    image_url = serializers.ImageField(required=False)

    class Meta:
        model = Profesor
        fields = ["jmbg", "ime", "prezime", "image_url"]