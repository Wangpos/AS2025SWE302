describe("Simple interaction test", () => {
  it("visits and interacts with the page", () => {
    cy.visit("/"); // navigate to baseUrl or full URL
    cy.get(".button").contains("Click me").click(); // find button by class + text and click
    cy.get('input[type="text"]').type("Hello"); // type into an input
  });

  it("asserts visibility, text and class (should + expect examples)", () => {
    // visit the page (uses cypress.config baseUrl if set)
    cy.visit("/");

    // Should-style assertions (most common)
    cy.get(".title").should("be.visible");
    cy.get(".title").should("have.text", "Welcome");
    cy.get(".title").should("have.class", "active");

    // Expect-style assertions (BDD style) - get the element then assert
    cy.get(".title").then(($el) => {
      expect($el).to.be.visible;
      expect($el.text()).to.equal("Welcome");
    });
  });
});
