script no server do graphql

mutation createCategory{
    createCategory(input:{name:"Tech", description:"Cursos Tech"}){
        id
        name
        description
    }
}

mutation createCourse{
    createCourse(input:{name: "fullcycle", description: "the best", categoryId:"bbc029ea-61ec-4f1f-818f-7874fc28249f"}){
        id
        name
    }
}

query queryCategories{
    categories{
        id
        name
        description
    }
}

query queryCategoriesWithCourses{
    categories{
        id
        name
        description
        courses{
            id
            name
        }
    }
}

query queryCouses{
    courses{
        id
        name
    }
}

query queryCousesWithCategory{
    courses{
        id
        name
        description
        category{
            id
            name
            description
        }
    }
}

para rodar o server
go run .\cmd\server\server.go

para gerar as modificações
go run github.com/99designs/gqlgen generate
