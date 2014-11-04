package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
    //"fmt"
)

const COLLECTION = "teams"


type Team struct {
    Id       bson.ObjectId `bson:"_id,omitempty"`
    Name     string        `bson:"name"`
    Users    []string      `bson:"users"`
    Projects []Project     `bson:"projects"`
}


type Project struct {
    Id      string    `bson:"id,omitempty"`
    Name    string    `bson:"name"`
    Users   []string  `bson:"users"`
    Todos   []Todo    `bson:"todos"`
}

type Todo struct {
    Id        bson.ObjectId `bson:"_id,omitempty"`
    AuthorId  string        `bson:"author_id"`
    Title     string        `bson:"title"`
    CDate     string        `bson:"c_date"`
    SSDate    string        `bson:"s_s_date"`
    SFDate    string        `bson:"s_f_date"`
    FDate     string        `bson:"f_date"`
    Finished  string        `bson:"finished"`
}



type TeamList struct {
    Id    bson.ObjectId  `bson:"_id,omitempty"`
	Name  string         `bson:"name"`
}

type ProjectList struct {
    Id        bson.ObjectId      `bson:"_id,omitempty"`
    Name      string             `bson:"name"`
	Users     []string           `bson:"users"`
	Projects  []ProjectListItem  `bson:"projects"`
}

type ProjectListItem struct {
    Id    string  `bson:"id"`
    Name  string  `bson:"name"`
}



type TeamListResponse struct {
    Id    string
    Name  string
}

type ProjectListResponse struct {
    Id    string
    Name  string
}



func Collection(d *mgo.Database) *mgo.Collection {
	return d.C(COLLECTION)
}

// get team list (API_G001)
func FindTeamList(d *mgo.Database) []TeamListResponse {
    teamList := []TeamList{}
    result   := []TeamListResponse{}
    Collection(d).Find(nil).All(&teamList)

    // TeamList to TeamListResponse
    for _, each := range teamList {
        newTeamList := TeamListResponse{}
        newTeamList.Id = each.Id.Hex()
        newTeamList.Name = each.Name
        result = append(result, newTeamList)
    }

    return result
}


// get project list (API_P001)
func FindProjectListById(d *mgo.Database, HexId string) []ProjectListResponse {
    projectList := []ProjectList{}
    result      := []ProjectListResponse{}
    if bson.IsObjectIdHex(HexId) {
        Id := bson.ObjectIdHex(HexId)
        Collection(d).FindId(Id).All(&projectList)
    }

    for _, each := range projectList {
        newPList := ProjectListResponse{}
        newPList.Id = each.Id.Hex()
        newPList.Name = each.Name
        result = append(result, newPList)
    }

    return result
}

// get todo list (API_T001)
func FindTodoListByTeamAndProjectId(d *mgo.Database, TeamHexId string, ProjectId string) []Todo {
    pResult := Team{}
    result  := []Todo{}
    if bson.IsObjectIdHex(TeamHexId) {
        Id := bson.ObjectIdHex(TeamHexId)
        Collection(d).FindId(Id).One(&pResult)

        for _, Project := range pResult.Projects {
            if Project.Id == ProjectId {
                result = Project.Todos
            }
        }

    }
    return result
}

