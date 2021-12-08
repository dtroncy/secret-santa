$(document).ready(function(){            
	var $fieldCount = 2;
	$('#addField').on('click', function(){
		$fieldCount++;;
		$( "#participants" ).append( "<tr class=\"participant\"><td><label for=\"name" + $fieldCount + "\">Participant's name n°" + $fieldCount + " : </label><input type=\"text\" name=\"name" + $fieldCount + "\" id=\"name" + $fieldCount + "\" required><td><label for=\"email" + $fieldCount + "\">Email : </label><input type=\"email\" name=\"email" + $fieldCount + "\" id=\"email" + $fieldCount + "\" required></td><td><span class=\"removeField\">Remove</span></td></tr>" );
	});
	
	$('form').on('click', '.removeField', function(){
		$(this).parent().parent().remove();
	});
});