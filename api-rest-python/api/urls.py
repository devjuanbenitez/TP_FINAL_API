from django.urls import path
from .views import MarcaListCreateView, MarcaDetailView, AllMarcasListView, MarcaDeleteView, MarcaUpdateView, ProductoListCreateView, ProductoDetailView, AllProductosListView, ProductoDeleteView, ProductoUpdateView, CategoriaListCreateView, CategoriaDetailView,AllCategoriasListView,CategoriaDeleteView,CategoriaUpdateView


urlpatterns = [
    path('marcas', MarcaListCreateView.as_view(), name='marca-list-create'),
    path('marcas/<int:pk>', MarcaDetailView.as_view(), name='marca-detail'),
    path('marcas', AllMarcasListView.as_view(), name='all-marcas-list'),  
    path('marcas/<int:pk>', MarcaDeleteView.as_view(), name='marca-delete'), 
    path('marcas/<int:pk>', MarcaUpdateView.as_view(), name='marca-update'), 
    path('productos', ProductoListCreateView.as_view(), name='producto-list-create'),
    path('productos/<int:pk>', ProductoDetailView.as_view(), name='producto-detail'),
    path('productos', AllProductosListView.as_view(), name='all-productos-list'),  
    path('productos/<int:pk>', ProductoDeleteView.as_view(), name='producto-delete'), 
    path('productos/<int:pk>', ProductoUpdateView.as_view(), name='producto-update'), 
    path('categorias', CategoriaListCreateView.as_view(), name='categoria-list-create'),
    path('categorias/<int:pk>', CategoriaDetailView.as_view(), name='categoria-detail'),
    path('categorias', AllCategoriasListView.as_view(), name='all-categorias-list'),  
    path('categorias<int:pk>', CategoriaDeleteView.as_view(), name='categoria-delete'), 
    path('categorias/<int:pk>', CategoriaUpdateView.as_view(), name='categoria-update'), 
]