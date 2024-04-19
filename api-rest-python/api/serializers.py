from rest_framework import serializers
from .models import Producto, Categoria, Marca

from rest_framework import serializers
from .models import Producto, Categoria, Marca


class CategoriaSerializer(serializers.ModelSerializer):
    class Meta:
        model = Categoria
        fields = '__all__'
        
class MarcaSerializer(serializers.ModelSerializer):
   
    def validate_id(self, value):
        """
        Validar que la marca con el ID especificado exista.
        """
        if not Marca.objects.filter(id=value).exists():
            raise serializers.ValidationError(f"La marca con ID {value} no existe.")
        return value  
    class Meta:
        model = Marca
        fields = '__all__'


class ProductoSerializer(serializers.ModelSerializer):
       
    marca = serializers.PrimaryKeyRelatedField(
        queryset=Marca.objects.all(),
        required=True,
        error_messages={'does_not_exist': 'La marca con ID {pk_value} no existe.'}
    )

    categoria = serializers.PrimaryKeyRelatedField(
        queryset=Categoria.objects.all(),
        required=True,
        error_messages={'does_not_exist': 'La categoría con ID {pk_value} no existe.'}
    )

    class Meta:
        model = Producto
        fields = '__all__'