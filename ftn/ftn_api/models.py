from django.db import models
# from django.contrib.auth.models import User

def upload_to(instance, filename):
    return 'images/{filename}'.format(filename=filename)

class Student(models.Model):
    jmbg = models.CharField(max_length = 180, unique=True)
    ime = models.CharField(max_length = 180)
    prezime = models.CharField(max_length = 180)
    image_url = models.ImageField(upload_to=upload_to, blank=True, null=True)


    def __str__(self):
        return self.jmbg


class Profesor(models.Model):
    jmbg = models.CharField(max_length = 180, unique=True)
    ime = models.CharField(max_length = 180)
    prezime = models.CharField(max_length = 180)
    image_url = models.ImageField(upload_to=upload_to, blank=True, null=True)


    def __str__(self):
        return self.jmbg