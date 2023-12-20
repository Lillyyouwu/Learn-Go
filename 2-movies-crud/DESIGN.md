### a movie server without database

            ROUTES      FUNCTIONS   ENDPOINTS   METHODS
            GET ALL     getMovies   /movies     GET
            GET BY ID   getMovie    /movies/id  GET
GORILLA     CREATE      createMovie /movies     POST
 MUX        UPDATE      updateMovie /movies/id  PUT
            DELETE      deleteMovie /movies/id  DELETE