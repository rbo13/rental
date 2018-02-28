$(document).ready(function() {
  $('select').material_select();
  $('.slider').slider();
  $(".button-collapse").sideNav();
  $('.collapsible').collapsible();
  $('.modal').modal({
    complete: function(modal) {
      $(modal).find('form').trigger('reset');
    }
  });

  $('#startDate').pickadate({
    selectMonths: true, // Creates a dropdown to control month
    selectYears: 101, // Creates a dropdown of 15 years to control year,
    today: 'Today',
    clear: 'Clear',
    close: 'Ok',
    closeOnSelect: false // Close upon selecting a date,
  });

  $('#endDate').pickadate({
    selectMonths: true,
    selectYears: 15,
    today: false,
    closeOnSelect: false
  });

  $('#birthDate').pickadate({
    selectMonths: true,
    selectYears: 15,
    today: false,
    closeOnSelect: false
  });

  $('.dropdown-button').dropdown({
    inDuration: 300,
    outDuration: 225,
    constrainWidth: true,
    hover: false,
    gutter: 0,
    belowOrigin: true,
    alignment: 'left',
    stopPropagation: false,
  });

  $('.dropdown-menu').dropdown({
    inDuration: 300,
    outDuration: 225,
    constrainWidth: false,
    hover: true,
    gutter: 0,
    belowOrigin: true,
    alignment: 'left',
    stopPropagation: false,
  });

  $('.carousel.carousel-slider').carousel({
    fullWidth: true
  });

  autoplay();

  function autoplay() {
    $('.carousel').carousel('next');
    setTimeout(autoplay, 4500);
  }
});