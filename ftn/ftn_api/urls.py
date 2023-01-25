from django.urls import path, include, re_path as url
from django.conf import settings
from django.conf.urls.static import static

from .views import (
    StudentApiView,
    ProfesorApiView,
)

urlpatterns = [
    path('student/api', StudentApiView.as_view()),
    path('profesor/api', ProfesorApiView.as_view()),
] + static(settings.MEDIA_URL, document_root=settings.MEDIA_ROOT)