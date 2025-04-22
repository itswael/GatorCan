describe("ProtectedDashboard Tests", () => {
  beforeEach(() => {
    cy.clearLocalStorage();
    cy.visit("http://localhost:5173");
  });

  it("Redirects to login if no user is logged in", () => {
    cy.visit("http://localhost:5173/dashboard");
    cy.url().should("include", "/login");
  });

  it("Redirects to dashboard if user lacks required roles", () => {
    cy.visit("http://localhost:5173/dashboard", {
      onBeforeLoad: (win) => {
        win.localStorage.setItem("username", "testuser");
        win.localStorage.setItem("roles", JSON.stringify(["guest"]));
      },
    });
    cy.url().should("include", "/dashboard");
  });

  it("Allows access if user has the correct role", () => {
    cy.visit("http://localhost:5173/dashboard", {
      onBeforeLoad: (win) => {
        win.localStorage.setItem("username", "testuser");
        win.localStorage.setItem("roles", JSON.stringify(["admin"]));
      },
    });

    cy.contains("GATORCAN-ADMIN").should("be.visible");
  });
});
