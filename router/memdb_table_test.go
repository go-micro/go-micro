package router

import "testing"

func testMemDBTableSetup() (*memDBTable, Route) {
	table := NewMemDBTable()

	route := Route{
		Service: "dest.svc",
		Address: "dest.addr",
		Gateway: "dest.gw",
		Network: "dest.network",
		Router:  "src.router",
		Link:    "det.link",
		Metric:  10,
	}

	return table, route
}

func TestMemDBTable_Create(t *testing.T) {
	table, route := testMemDBTableSetup()

	if err := table.Create(route); err != nil {
		t.Errorf("error adding route: %s", err)
	}

	// adds new route for the original destination
	route.Gateway = "dest.gw2"

	if err := table.Create(route); err != nil {
		t.Errorf("error adding route: %s", err)
	}

	// adding the same route under Insert policy must error
	if err := table.Create(route); err != ErrDuplicateRoute {
		t.Errorf("error adding route. Expected error: %s, found: %s", ErrDuplicateRoute, err)
	}
}

func TestMemDBTable_Delete(t *testing.T) {
	table, route := testMemDBTableSetup()

	if err := table.Create(route); err != nil {
		t.Errorf("error adding route: %s", err)
	}

	// should fail to delete non-existant route
	prevSvc := route.Service
	route.Service = "randDest"

	if err := table.Delete(route); err != ErrRouteNotFound {
		t.Errorf("error deleting route. Expected: %s, found: %s", ErrRouteNotFound, err)
	}

	// we should be able to delete the existing route
	route.Service = prevSvc

	if err := table.Delete(route); err != nil {
		t.Errorf("error deleting route: %s", err)
	}
}

func TestMemDBTable_Update(t *testing.T) {
	table, route := testMemDBTableSetup()

	if err := table.Create(route); err != nil {
		t.Errorf("error adding route: %s", err)
	}

	// change the metric of the original route
	route.Metric = 200

	if err := table.Update(route); err != nil {
		t.Errorf("error updating route: %s", err)
	}

	// this should add a new route
	route.Service = "rand.dest"

	if err := table.Update(route); err != nil {
		t.Errorf("error updating route: %s", err)
	}
}
func TestMemDBTable_List(t *testing.T) {
	table, route := testMemDBTableSetup()

	svc := []string{"one.svc", "two.svc", "three.svc"}

	for i := 0; i < len(svc); i++ {
		route.Service = svc[i]
		if err := table.Create(route); err != nil {
			t.Errorf("error adding route: %s", err)
		}
	}

	routes, err := table.List()
	if err != nil {
		t.Errorf("error listing routes: %s", err)
	}

	if len(routes) != len(svc) {
		t.Errorf("incorrect number of routes listed. Expected: %d, found: %d", len(svc), len(routes))
	}
}

//func TestMemDBTable_Query(t *testing.T) {
//	table, route := testMemDBTableSetup()
//
//	svc := []string{"svc1", "svc2", "svc3"}
//	net := []string{"net1", "net2", "net1"}
//	gw := []string{"gw1", "gw2", "gw3"}
//	rtr := []string{"rtr1", "rt2", "rt3"}
//
//	for i := 0; i < len(svc); i++ {
//		route.Service = svc[i]
//		route.Network = net[i]
//		route.Gateway = gw[i]
//		route.Router = rtr[i]
//		if err := table.Create(route); err != nil {
//			t.Errorf("error adding route: %s", err)
//		}
//	}
//
//	// return all routes
//	routes, err := table.Query()
//	if err != nil {
//		t.Errorf("error looking up routes: %s", err)
//	}
//
//	// query routes particular network
//	network := "net1"
//
//	routes, err = table.Query(QueryNetwork(network))
//	if err != nil {
//		t.Errorf("error looking up routes: %s", err)
//	}
//
//	if len(routes) != 2 {
//		t.Errorf("incorrect number of routes returned. Expected: %d, found: %d", 2, len(routes))
//	}
//
//	for _, route := range routes {
//		if route.Network != network {
//			t.Errorf("incorrect route returned. Expected network: %s, found: %s", network, route.Network)
//		}
//	}
//
//	// query routes for particular gateway
//	gateway := "gw1"
//
//	routes, err = table.Query(QueryGateway(gateway))
//	if err != nil {
//		t.Errorf("error looking up routes: %s", err)
//	}
//
//	if len(routes) != 1 {
//		t.Errorf("incorrect number of routes returned. Expected: %d, found: %d", 1, len(routes))
//	}
//
//	if routes[0].Gateway != gateway {
//		t.Errorf("incorrect route returned. Expected gateway: %s, found: %s", gateway, routes[0].Gateway)
//	}
//
//	// query routes for particular router
//	router := "rtr1"
//
//	routes, err = table.Query(QueryRouter(router))
//	if err != nil {
//		t.Errorf("error looking up routes: %s", err)
//	}
//
//	if len(routes) != 1 {
//		t.Errorf("incorrect number of routes returned. Expected: %d, found: %d", 1, len(routes))
//	}
//
//	if routes[0].Router != router {
//		t.Errorf("incorrect route returned. Expected router: %s, found: %s", router, routes[0].Router)
//	}
//
//	// query particular gateway and network
//	query := []QueryOption{
//		QueryGateway(gateway),
//		QueryNetwork(network),
//		QueryRouter(router),
//	}
//
//	routes, err = table.Query(query...)
//	if err != nil {
//		t.Errorf("error looking up routes: %s", err)
//	}
//
//	if len(routes) != 1 {
//		t.Errorf("incorrect number of routes returned. Expected: %d, found: %d", 1, len(routes))
//	}
//
//	if routes[0].Gateway != gateway {
//		t.Errorf("incorrect route returned. Expected gateway: %s, found: %s", gateway, routes[0].Gateway)
//	}
//
//	if routes[0].Network != network {
//		t.Errorf("incorrect network returned. Expected network: %s, found: %s", network, routes[0].Network)
//	}
//
//	if routes[0].Router != router {
//		t.Errorf("incorrect route returned. Expected router: %s, found: %s", router, routes[0].Router)
//	}
//
//	// non-existen route query
//	routes, err = table.Query(QueryService("foobar"))
//	if err != ErrRouteNotFound {
//		t.Errorf("error looking up routes. Expected: %s, found: %s", ErrRouteNotFound, err)
//	}
//
//	if len(routes) != 0 {
//		t.Errorf("incorrect number of routes returned. Expected: %d, found: %d", 0, len(routes))
//	}
//}
