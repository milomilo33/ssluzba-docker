from django.urls import path, include, re_path as url

from .views import (
    StudentApiView,
    ProfesorApiView,
)

urlpatterns = [
    path('student/api', StudentApiView.as_view()),
    path('profesor/api', ProfesorApiView.as_view()),
]