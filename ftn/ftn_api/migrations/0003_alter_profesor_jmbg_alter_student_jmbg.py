# Generated by Django 4.1.5 on 2023-01-26 02:42

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('ftn_api', '0002_profesor_image_url_student_image_url'),
    ]

    operations = [
        migrations.AlterField(
            model_name='profesor',
            name='jmbg',
            field=models.CharField(max_length=180),
        ),
        migrations.AlterField(
            model_name='student',
            name='jmbg',
            field=models.CharField(max_length=180),
        ),
    ]
