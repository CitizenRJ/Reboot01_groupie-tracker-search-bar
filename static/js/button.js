console.log("tomfifugyuiehj")
$(document).ready(function () {
  $(".checkboxes").click(function () {
    $("#contactChoiceAll").prop("checked", false);
  });
  $("#contactChoiceAll").click(function () {
    $(".checkboxes").prop("checked", false);
  });
});

